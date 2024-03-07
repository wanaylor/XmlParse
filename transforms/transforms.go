package transforms

import (
    "fmt"
    "os"
    "encoding/json"
)

type transformer interface {
    GetMap() transformMap
}

type transformMap struct {
    Level1 map[string] string
    Level2 map[string] map[string] string
}

type JsonTransformer struct {
    oldToNewRelation transformMap
}

func (j *JsonTransformer) GetMap() transformMap {
    return j.oldToNewRelation
}

func (j *JsonTransformer) Read(inputFile string) {
    v, err := os.ReadFile(inputFile)
    if err != nil{
        panic(fmt.Sprintf("Could not read file %s", inputFile))
    }
}
