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
	source_start int
	source_end   int
}

type seed struct {
	id          int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

func Run() {
	file := utils.GetData("5", true)
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n\n")

	defer utils.Timer("day5")()

	seeds := utils.ConvertStringSliceToIntSlice(strings.Split(lines[0], " ")[1:])
	Part1(seeds, lines)

}

func Part1(seeds []int, lines []string) {
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

	seedToSoilMap := sourceToDestination(mapTypes[SOIL])
	soilToFertMap := sourceToDestination(mapTypes[FERTILIZER])
	fertToWaterMap := sourceToDestination(mapTypes[WATER])
	waterToLight := sourceToDestination(mapTypes[LIGHT])
	lightToTemp := sourceToDestination(mapTypes[TEMPERATURE])
	tempToHumid := sourceToDestination(mapTypes[HUMIDITY])
	humidToLocation := sourceToDestination(mapTypes[LOCATION])

	specifiedSeeds := []seed{}
	for _, s := range seeds {
		newSeed := seed{id: s}

		newSeed.soil = AddSpecsToSeed(newSeed.id, seedToSoilMap)
		newSeed.fertilizer = AddSpecsToSeed(newSeed.soil, soilToFertMap)
		newSeed.water = AddSpecsToSeed(newSeed.fertilizer, fertToWaterMap)
		newSeed.light = AddSpecsToSeed(newSeed.water, waterToLight)
		newSeed.temperature = AddSpecsToSeed(newSeed.light, lightToTemp)
		newSeed.humidity = AddSpecsToSeed(newSeed.temperature, tempToHumid)
		newSeed.location = AddSpecsToSeed(newSeed.humidity, humidToLocation)

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

func AddSpecsToSeed(sourceId int, ranges []mapRange) int {

	for _, mapRange := range ranges {
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
