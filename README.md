# RankVectors Go SDK

Official Go SDK for the RankVectors API - Intelligent internal linking optimization using AI.

## Installation

```bash
go get github.com/rankvectors/rankvectors-go-sdk
```

## Quick Start

```go
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
```

## Features

- ✅ **Go 1.19+**: Modern Go support with modules
- ✅ **Context Support**: Full context support for cancellation and timeouts
- ✅ **Type Safety**: Strongly typed models and request/response objects
- ✅ **Error Handling**: Comprehensive error handling with detailed error messages
- ✅ **HTTP Client**: Configurable HTTP client with retry logic
- ✅ **Concurrent Safe**: Thread-safe API client for concurrent usage

## API Coverage

- **Projects**: Create, manage, and delete projects
- **Crawling**: Start and monitor website crawls
- **Suggestions**: Generate AI-powered link suggestions
- **Implementations**: Implement links via API or webhook
- **Credits**: Manage credit balance and usage
- **Content Verification**: Verify content changes before implementation

## Error Handling

```go
project, response, err := apiClient.ProjectsAPI.CreateProject(ctx).CreateProjectRequest(projectData).Execute()
if err != nil {
    if response != nil {
        switch response.StatusCode {
        case 401:
            fmt.Println("Authentication failed")
        case 400:
            fmt.Printf("Invalid request: %s\n", response.Body)
        default:
            fmt.Printf("Unexpected error: %v\n", err)
        }
    } else {
        fmt.Printf("Network error: %v\n", err)
    }
    return
}
```

## Context Usage

```go
// Create context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

// Use context in API calls
project, _, err := apiClient.ProjectsAPI.CreateProject(ctx).CreateProjectRequest(projectData).Execute()
if err != nil {
    if ctx.Err() == context.DeadlineExceeded {
        fmt.Println("Request timed out")
    } else {
        fmt.Printf("Error: %v\n", err)
    }
    return
}
```

## Configuration

```go
configuration := rankvectors.NewConfiguration()
configuration.Servers = []rankvectors.ServerConfiguration{
    {
        URL: "https://api.rankvectors.com",
    },
}
configuration.DefaultHeader["Authorization"] = "Bearer YOUR_API_KEY"
configuration.Timeout = 30 * time.Second
configuration.UserAgent = "my-app/1.0.0"
```

## Concurrent Usage

```go
func processProjects(apiClient *rankvectors.APIClient, projectIDs []string) {
    var wg sync.WaitGroup
    results := make(chan *projects.Project, len(projectIDs))
    
    for _, id := range projectIDs {
        wg.Add(1)
        go func(projectID string) {
            defer wg.Done()
            
            project, _, err := apiClient.ProjectsAPI.GetProject(context.Background(), projectID).Execute()
            if err != nil {
                fmt.Printf("Error getting project %s: %v\n", projectID, err)
                return
            }
            
            results <- project
        }(id)
    }
    
    wg.Wait()
    close(results)
    
    for project := range results {
        fmt.Printf("Project: %s\n", *project.Name)
    }
}
```

## Support

- **Documentation**: https://rankvectors.com/docs
- **Support**: tj@rankvectors.com
- **GitHub**: https://github.com/RankVectors/go-sdk

## License

MIT License