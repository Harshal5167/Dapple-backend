package utils 

import (
	"strings"	
	"fmt"
)

func BuildStringForLevels(levelDetails []map[string]string) string {
    var builder strings.Builder
    
    builder.WriteString("Available Levels:\n")
    
    for i, level := range levelDetails {
        builder.WriteString(fmt.Sprintf(
            "Level %d:\n"+
            "- ID: %s\n"+
            "- Name: %s\n"+
            "- Description: %s\n\n",
            i+1,
            level["levelId"],
            level["levelName"],
            level["description"],
        ))
    }
    
    return builder.String()
}