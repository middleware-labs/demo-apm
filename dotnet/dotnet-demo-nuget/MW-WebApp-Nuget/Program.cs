using Middleware.APM;

var builder = WebApplication.CreateBuilder(args);

var mwApiKey = builder.Configuration.GetSection("MW")["ApiKey"];
var target = builder.Configuration.GetSection("MW")["Target_URL"];

var attributes = new Dictionary<string, object>
{
    { "mw.account_key", mwApiKey },
    { "runtime.metrics.dotnet", true },
    { "project.name", "demo-apm" },
    { "service.name", "dotnet-webapp-nuget" },
    { "target", target },
    { "console.exporter", true }
};

// builder.Services.ConfigureMWInstrumentation(attributes);

// Add services to the container.
var configuration = new ConfigurationBuilder()
    .SetBasePath(Directory.GetCurrentDirectory())
    .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
    .AddEnvironmentVariables()
    .Build();

builder.Services.ConfigureMWInstrumentation(configuration);

builder.Logging.AddConfiguration(configuration.GetSection("Logging"));
builder.Logging.AddConsole();
builder.Logging.SetMinimumLevel(LogLevel.Debug);

var app = builder.Build();

Logger.Init(app.Services.GetRequiredService<ILoggerFactory>());

// Configure the HTTP request pipeline.
app.UseSwagger();
app.UseSwaggerUI();

app.UseHttpsRedirection();

// app.UseAuthorization();

app.MapControllers();

app.Run();
