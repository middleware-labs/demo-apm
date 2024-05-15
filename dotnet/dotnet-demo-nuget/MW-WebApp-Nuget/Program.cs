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

builder.Services.ConfigureMWInstrumentation(attributes);

// Add services to the container.

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

Logger.Init(app.Services.GetRequiredService<ILoggerFactory>());

// Configure the HTTP request pipeline.
app.UseSwagger();
app.UseSwaggerUI();

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
