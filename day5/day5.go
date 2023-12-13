package day5

import (
	"fmt"
	"main/utils"
	"math"
	"slices"
	"strings"
	"sync"
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

func Run() {
	file := utils.GetData("5", true)
	file = strings.Trim(file, "\n")
	lines := strings.Split(file, "\n\n")

	defer utils.Timer("day5")()

	seeds := utils.ConvertStringSliceToIntSlice(strings.Split(lines[0], " ")[1:])

	part1Seeds := []mapRange{}
	for _, s := range seeds {
		part1Seeds = append(part1Seeds, mapRange{source_start: s, source_end: s + 1})
	}

	solve(part1Seeds, lines)

	part2Seeds := seedRange(seeds)
	solve(part2Seeds, lines)

}

func solve(seeds []mapRange, lines []string) {

	mapTypes := getMapTypes(lines)
	maps := map[string][]mapRange{}

	maps[SOIL] = sourceToDestination(mapTypes[SOIL])
	maps[FERTILIZER] = sourceToDestination(mapTypes[FERTILIZER])
	maps[WATER] = sourceToDestination(mapTypes[WATER])
	maps[LIGHT] = sourceToDestination(mapTypes[LIGHT])
	maps[TEMPERATURE] = sourceToDestination(mapTypes[TEMPERATURE])
	maps[HUMIDITY] = sourceToDestination(mapTypes[HUMIDITY])
	maps[LOCATION] = sourceToDestination(mapTypes[LOCATION])

	lowestLocation := 0
	max := 100

	for _, seedRange := range seeds {
		var wg sync.WaitGroup
		jobQueue := make(chan int, max)

		for j := 0; j < max; j++ {
			wg.Add(1)

			go func(ch chan int) {
				defer wg.Done()

				for seedId := range ch {
					soil := addSpecsToSeed(seedId, maps[SOIL])
					fertilizer := addSpecsToSeed(soil, maps[FERTILIZER])
					water := addSpecsToSeed(fertilizer, maps[WATER])
					light := addSpecsToSeed(water, maps[LIGHT])
					temperature := addSpecsToSeed(light, maps[TEMPERATURE])
					humidity := addSpecsToSeed(temperature, maps[HUMIDITY])
					location := addSpecsToSeed(humidity, maps[LOCATION])

					if lowestLocation == 0 || location < lowestLocation {
						lowestLocation = location
					}
				}

			}(jobQueue)
		}
		for i := seedRange.source_start; i < seedRange.source_end; i++ {
			jobQueue <- i
		}

		close(jobQueue)
		wg.Wait()
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

func addSpecsToSeed(sourceId int, ranges []mapRange) int {

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
