using System;
using System.Net.Http;
using System.Threading.Tasks;

internal class Program
{
    private static async Task Main(string[] args)
    {
        Console.WriteLine("Hello World from async Main!");

        // Use the first command-line argument as the URL or default to a preset URL.
        string url = args.Length > 0 ? args[0] : "https://example.com";
        Console.WriteLine($"Fetching content from: {url}");

        using HttpClient client = new HttpClient();
        try
        {
            // Asynchronously fetch the content of the URL.
            string content = await client.GetStringAsync(url);
            Console.WriteLine("Fetched content (first 200 characters):");
            Console.WriteLine(content.Substring(0, Math.Min(200, content.Length)));
        }
        catch (Exception ex)
        {
            Console.WriteLine("An error occurred: " + ex.Message);
        }

        // Web application
        // var configuration = new ConfigurationBuilder()
        //     .SetBasePath(Directory.GetCurrentDirectory())
        //     .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
        //     .AddEnvironmentVariables()
        //     .Build();

        // var builder = WebApplication.CreateBuilder(args);
        // builder.Services.ConfigureMWInstrumentation(configuration);

        // builder.Logging.AddConfiguration(configuration.GetSection("Logging"));
        // builder.Logging.AddConsole();
        // builder.Logging.SetMinimumLevel(LogLevel.Debug);

        // // Add the following code after var app = builder.Build();
        // Logger.Init(app.Services.GetRequiredService<ILoggerFactory>());

        // Build configuration from appsettings.json and environment variables
        var configuration = new ConfigurationBuilder()
            .SetBasePath(Directory.GetCurrentDirectory())
            .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
            .AddEnvironmentVariables()
            .Build();

        // Create a Generic Host (works in console apps) and register services
        var host = Host.CreateDefaultBuilder(args)
            .ConfigureServices((context, services) =>
            {
                // Replace or add your custom instrumentation configuration
                services.ConfigureMWInstrumentation(configuration);
                // Register other services if needed...
            })
            .ConfigureLogging(logging =>
            {
                logging.ClearProviders();
                logging.AddConfiguration(configuration.GetSection("Logging"));
                logging.AddConsole();
                logging.SetMinimumLevel(LogLevel.Debug);
            })
            .Build();

        // After building the host, initialize your custom logger with ILoggerFactory
        Logger.Init(host.Services.GetRequiredService<ILoggerFactory>());



        // Wait for user input before closing.
        Console.WriteLine("Press any key to exit...");
        Console.ReadKey();
    }
}
