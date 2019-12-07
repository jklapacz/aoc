package day06

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type OrbitMap struct {
	Orbits map[string]*Orbitter
}

type Orbitter struct {
	id string
	directOrbitCount, indirectOrbitCount int
	indirect []*Orbitter
	parent *Orbitter
}

func (o *Orbitter) ToString() string {
	return fmt.Sprintf("id: [%v]\n\tdirect=%v\n\tindirect=%v\n\tparent=%v", o.id, o.directOrbitCount, len(o.indirect), o.parent)
}

func buildIndirect(o *Orbitter) []*Orbitter {
	if o == nil {
		return nil
	}
	return append(buildIndirect(o.parent), o)
}

func buildIndirectList(orbitter *Orbitter, com *Orbitter) {
	if com == nil || orbitter == nil {
		return
	}
	//if orbitter.parent.id != com.id {
	orbitter.indirect = append(orbitter.indirect, com)
	//}
	buildIndirectList(orbitter, com.parent)
}

func ExtractCOMObjects(orbits *OrbitMap, gravityMapping string) {
	objects := strings.Split(gravityMapping, ")")
	if len(objects) != 2 {
		log.Printf("bad input! %v\n", gravityMapping)
		return
	}
	idCom := objects[0]
	idOrbitter := objects[1]
	var com, orbitter *Orbitter
	if _, ok := orbits.Orbits[idCom]; !ok {
		com = &Orbitter{
			id:                 idCom,
			directOrbitCount:   0,
			indirectOrbitCount: 0,
			indirect: []*Orbitter{},
			parent:             nil,
		}
		orbits.Orbits[com.id] = com
	} else {
		com = orbits.Orbits[idCom]
	}
	if _, ok := orbits.Orbits[idOrbitter]; !ok {
		orbitter = &Orbitter{
			id:                	idOrbitter,
			directOrbitCount:   com.directOrbitCount + 1,
			indirectOrbitCount: com.directOrbitCount,
			indirect: 			[]*Orbitter{},
			parent:             com,
		}
		orbitter.indirect = buildIndirect(orbitter.parent)
		orbits.Orbits[orbitter.id] = orbitter
	} else {
		orbitter = orbits.Orbits[idOrbitter]
		orbitter.directOrbitCount= com.directOrbitCount + 1
		orbitter.indirectOrbitCount = com.directOrbitCount
		orbitter.parent = com
		//orbitter.indirect = buildIndirect(orbitter.parent)
		orbits.Orbits[orbitter.id] = orbitter
	}
	//buildIndirectList(orbitter, com)
}

func (o *OrbitMap) Parse(gravityMapping string) {
	ExtractCOMObjects(o, gravityMapping)
}

func (o *OrbitMap) totalOrbits() int {
	if len(o.Orbits) == 0 {
		return 0
	}
	sum := 0
	for _, com := range o.Orbits {
		sum += len(com.indirect)
		//sum += com.directOrbitCount
	}
	return sum
}

func (o *OrbitMap) ToString() string {
	outputString := ""
	for _, orbitter := range o.Orbits {
		outputString += fmt.Sprintf("%v\n", orbitter.ToString())
	}
	return outputString
}

func Solve(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	orbitMap := &OrbitMap{Orbits:make(map[string]*Orbitter)}
	for scanner.Scan() {
		orbitMap.Parse(scanner.Text())
	}
	fmt.Println(orbitMap.ToString())

	return orbitMap.totalOrbits()
}