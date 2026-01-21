package registery

import (
	"github.com/muskiteer/GoCP/server/structs"
	"github.com/muskiteer/GoCP/server/tools"
	"encoding/json"
	"os"
	"fmt"
	"context"
)

type Registry struct {
    Tools map[string]structs.Tool
}

var executorMap = map[string]structs.ToolExecutor{
    "fetching_crypto": tools.FetchCryptoData,
    "fetching_wikipedia": tools.FetchWikipediaData,
}

func LoadToolManifest(path string) (*structs.ToolManifest, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var manifest structs.ToolManifest
    if err := json.Unmarshal(data, &manifest); err != nil {
        return nil, fmt.Errorf("failed in LoadToolManifest: %v", err)
    }

    return &manifest, nil
}


func InitRegistry(manifest *structs.ToolManifest) (*Registry, error) {
    r := &Registry{
        Tools: make(map[string]structs.Tool),
    }

    for _, spec := range manifest.Tools {
        exec, ok := executorMap[spec.Name]
        if !ok {
            return nil, fmt.Errorf("no executor found for tool: %s", spec.Name)
        }

        r.Tools[spec.Name] = structs.Tool{
            Spec:    spec,
            Execute: exec,
        }
    }

    return r, nil
}

func (r *Registry) Execute(
    ctx context.Context,
    toolName string,
    args map[string]any,
) (any, error) {

    tool, ok := r.Tools[toolName]
    if !ok {
        return nil, fmt.Errorf("tool not found: %s", toolName)
    }

   
    return tool.Execute(ctx, args)
}




