package common

type Record struct {
        Time int64 `json:"time"`
        Prompt string `json:"prompt"`
        Output string `json:"output"`
}
