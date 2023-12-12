package day5

import (
	"fmt"
	"main/utils"
	"math"
	"slices"
	"strings"
)

const (
	SOIL        = "seed-to-soil"
	FERTILIZER  = "soil-to-fertilizer"
	WATER       = "fertilizer-to-water"
	LIGHT       = "water-to-light"
	TEMPERATURE = "light-to-temperature"
	HUMIDITY    = "temperature-to-humidity"
	LOCATION    = "humidity-to-location"
)

type mapInfo struct {
	destination_range int
	source_range      int
	lenght_range      int
}

type mapRange struct {
	des_start    int
	des_end      int
	source_start int
	source_end   int
}

type seed struct {
	id       int
	location int
}

func Run() {
	file := utils.GetData("5", false)
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n\n")

	defer utils.Timer("day5")()

	seeds := utils.ConvertStringSliceToIntSlice(strings.Split(lines[0], " ")[1:])
	part1(seeds, lines)

	seedRanges := seedRange(seeds)
	part2(seedRanges, lines)

}

func part2(seeds []mapRange, lines []string) {

	mapTypes := getMapTypes(lines)

	seedToSoilMap := sourceToDestination(mapTypes[SOIL])
	soilToFertMap := sourceToDestination(mapTypes[FERTILIZER])
	fertToWaterMap := sourceToDestination(mapTypes[WATER])
	waterToLight := sourceToDestination(mapTypes[LIGHT])
	lightToTemp := sourceToDestination(mapTypes[TEMPERATURE])
	tempToHumid := sourceToDestination(mapTypes[HUMIDITY])
	humidToLocation := sourceToDestination(mapTypes[LOCATION])

	diff := 0
	for _, seedRange := range seeds {
		d := seedRange.source_end - seedRange.source_start
		diff += d
	}

	lowestLocation := 0
	for _, seedRange := range seeds {
		for i := seedRange.source_start; i < seedRange.source_end; i++ {

			soil := addSpecsToSeed(i, &seedToSoilMap)
			fertilizer := addSpecsToSeed(soil, &soilToFertMap)
			water := addSpecsToSeed(fertilizer, &fertToWaterMap)
			light := addSpecsToSeed(water, &waterToLight)
			temperature := addSpecsToSeed(light, &lightToTemp)
			humidity := addSpecsToSeed(temperature, &tempToHumid)
			location := addSpecsToSeed(humidity, &humidToLocation)

			if lowestLocation == 0 || location < lowestLocation {
				lowestLocation = location
			}
		}
	}

	fmt.Println(lowestLocation)
}

func seedRange(seeds []int) []mapRange {

	seedRanges := []mapRange{}
	for i := 0; i < len(seeds); i += 2 {
		startSeed := seeds[i]
		endSeed := startSeed + seeds[i+1]
		seedRanges = append(seedRanges, mapRange{source_start: seeds[i], source_end: endSeed})

	}

	return seedRanges
}

func part1(seeds []int, lines []string) {

	mapTypes := getMapTypes(lines)

	seedToSoilMap := sourceToDestination(mapTypes[SOIL])
	soilToFertMap := sourceToDestination(mapTypes[FERTILIZER])
	fertToWaterMap := sourceToDestination(mapTypes[WATER])
	waterToLight := sourceToDestination(mapTypes[LIGHT])
	lightToTemp := sourceToDestination(mapTypes[TEMPERATURE])
	tempToHumid := sourceToDestination(mapTypes[HUMIDITY])
	humidToLocation := sourceToDestination(mapTypes[LOCATION])

	specifiedSeeds := make([]seed, 0, len(seeds))
	for _, s := range seeds {
		newSeed := seed{id: s}

		soil := addSpecsToSeed(newSeed.id, &seedToSoilMap)
		fertilizer := addSpecsToSeed(soil, &soilToFertMap)
		water := addSpecsToSeed(fertilizer, &fertToWaterMap)
		light := addSpecsToSeed(water, &waterToLight)
		temperature := addSpecsToSeed(light, &lightToTemp)
		humidity := addSpecsToSeed(temperature, &tempToHumid)
		newSeed.location = addSpecsToSeed(humidity, &humidToLocation)

		specifiedSeeds = append(specifiedSeeds, newSeed)
	}

	lowestLocation := 0
	for _, s := range specifiedSeeds {
		if lowestLocation == 0 || s.location < lowestLocation {
			lowestLocation = s.location
		}
	}

	fmt.Println(lowestLocation)
}

func getMapTypes(lines []string) map[string][]mapInfo {
	maps := lines[1:]

	mapTypes := map[string][]mapInfo{}

	for _, m := range maps {
		parsedMap := strings.Split(m, "\n")
		mapTitle := strings.Split(parsedMap[0], " ")[0]

		mapInfos := []mapInfo{}
		for i := 1; i < len(parsedMap); i++ {
			ranges := utils.ConvertStringSliceToIntSlice(strings.Split(parsedMap[i], " "))

			mapInfos = append(mapInfos, mapInfo{
				destination_range: ranges[0],
				source_range:      ranges[1],
				lenght_range:      ranges[2],
			})
		}

		mapTypes[mapTitle] = mapInfos
	}
	return mapTypes
}

func addSpecsToSeed(sourceId int, ranges *[]mapRange) int {

	for _, mapRange := range *ranges {
		if sourceId > mapRange.source_end || sourceId < mapRange.source_start {
			continue
		}

		diff := sourceId - mapRange.source_start
		return mapRange.des_start + diff
	}
	return sourceId
}

func sourceToDestination(ranges []mapInfo) []mapRange {

	slices.SortFunc(ranges, func(a, b mapInfo) int {
		return b.source_range - a.source_range
	})

	currentLimit := math.MaxInt
	sourceToDes := []mapRange{}
	for _, r := range ranges {

		end := r.source_range + r.lenght_range - 1
		mr := mapRange{
			des_start:    r.destination_range,
			des_end:      r.destination_range + r.lenght_range - 1,
			source_start: r.source_range,
			source_end:   end}

		if end > currentLimit && end < math.MaxInt {
			mr.source_end = currentLimit - 1
		}

		currentLimit = r.source_range
		sourceToDes = append(sourceToDes, mr)
	}

	slices.SortFunc(sourceToDes, func(a, b mapRange) int {
		return a.source_end - b.source_end
	})
	return sourceToDes
}
