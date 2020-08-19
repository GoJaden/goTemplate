package z

import "github.com/AlecAivazis/survey/v2"

type Answer struct {
	Project   string
	Namespace string
	Author    string
	Workdir   string
}

const (
	default_project   = "demo"
	default_author    = "x"
	default_namespace = "sdyxmall"
)

func NewAnswer() *Answer {
	CheckDir()
	return &Answer{
		Project:   default_project,
		Namespace: default_namespace,
		Author:    default_author,
		Workdir:   GetWorkDir(),
	}
}

func NewQ() []*survey.Question {
	return []*survey.Question{
		{
			Name: "Project",
			Prompt: &survey.Input{
				Message: "工程名: ",
				Default: default_project,
			},
			Validate: survey.Required,
		},
		{
			Name: "Namespace",
			Prompt: &survey.Select{
				Message: "命名空间: ",
				Options: []string{"sdyxmall", "mzmovie", "sdyxopen"},
				Default: default_namespace,
			},
			Validate: survey.Required,
		},
		{
			Name: "Author",
			Prompt: &survey.Input{
				Message: "作者名: ",
				Default: default_author,
			},
			Validate: survey.Required,
		},
	}
}
