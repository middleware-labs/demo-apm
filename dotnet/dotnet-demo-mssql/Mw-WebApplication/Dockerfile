FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base

WORKDIR /app
RUN apt-get update && apt-get install -y curl unzip

# Middleware
RUN curl -sSfL https://install.middleware.io/apm/dotnet/v1.4.1/scripts/mw-dotnet-auto-install.sh -O
RUN bash mw-dotnet-auto-install.sh 
RUN rm -rf mw-dotnet-auto-install.sh
RUN chmod +x $HOME/.mw-dotnet-auto/instrument.sh

# OpenTelemetry
# RUN curl -sSfL https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/releases/download/v1.4.0/otel-dotnet-auto-install.sh -O
# RUN sh ./otel-dotnet-auto-install.sh
# RUN rm -rf otel-dotnet-auto-install.sh
# RUN chmod +x $HOME/.otel-dotnet-auto/instrument.sh


FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY Mw-WebApplication.csproj .
RUN dotnet restore Mw-WebApplication.csproj
COPY . .
WORKDIR /src
RUN dotnet build -v normal -c Release -o /app/build > build_output.txt 2>&1

FROM build AS publish
RUN dotnet publish -c Release -o /app/publish 

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
COPY --from=build /src/build_output.txt .

# Middleware
CMD . $HOME/.mw-dotnet-auto/instrument.sh && dotnet Mw-WebApplication.dll

# OpenTelemetry
# CMD . $HOME/.otel-dotnet-auto/instrument.sh && dotnet Mw-WebApplication.dll
