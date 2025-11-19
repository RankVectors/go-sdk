package main

import (
    "context"
    "fmt"
    "os"
    
    "github.com/rankvectors/rankvectors-go-sdk"
    "github.com/rankvectors/rankvectors-go-sdk/projects"
)

func main() {
    // Configure the API client
    configuration := rankvectors.NewConfiguration()
    configuration.Servers = []rankvectors.ServerConfiguration{
        {
            URL: "https://api.rankvectors.com",
        },
    }
    
    // Set API key
    configuration.DefaultHeader["Authorization"] = "Bearer YOUR_API_KEY" // Replace with your actual API key
    
    // Create API client
    apiClient := rankvectors.NewAPIClient(configuration)
    
    // Create project
    projectData := projects.CreateProjectRequest{
        Name:   rankvectors.PtrString("My Website"),
        Domain: rankvectors.PtrString("https://example.com"),
        Prompt: rankvectors.PtrString("Only crawl blog posts and documentation"),
    }
    
    // Create project
    project, _, err := apiClient.ProjectsAPI.CreateProject(context.Background()).CreateProjectRequest(projectData).Execute()
    if err != nil {
        fmt.Printf("Error creating project: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Printf("Project created: %+v\n", project)
    
    // Start crawling
    crawl, _, err := apiClient.ProjectsAPI.StartCrawl(context.Background(), *project.Id).Execute()
    if err != nil {
        fmt.Printf("Error starting crawl: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Printf("Crawl started: %+v\n", crawl)
}
