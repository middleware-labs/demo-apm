FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base

WORKDIR /app
RUN apt-get update && apt-get install -y curl unzip

# Middleware
# RUN curl -sSfL https://install.middleware.io/apm/dotnet/v1.4.0/scripts/mw-dotnet-auto-install.sh -O
# RUN bash mw-dotnet-auto-install.sh 
# RUN rm -rf mw-dotnet-auto-install.sh
# RUN chmod +x $HOME/.mw-dotnet-auto/instrument.sh

# OpenTelemetry
# RUN curl -sSfL https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/releases/download/v1.4.0/otel-dotnet-auto-install.sh -O
# # COPY otel-dotnet-auto-install.sh otel-dotnet-auto-install.sh
# RUN sh ./otel-dotnet-auto-install.sh
# RUN chmod +x $HOME/.otel-dotnet-auto/instrument.sh


FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /src
COPY Mw-WebApplication.csproj .
RUN dotnet restore Mw-WebApplication.csproj
COPY . .
WORKDIR /src
RUN dotnet build -c Release -o /app/build

FROM build AS publish
RUN dotnet publish -c Release -o /app/publish 

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .

# Middleware
# CMD . $HOME/.mw-dotnet-auto/instrument.sh && dotnet Mw-WebApplication.dll

# OpenTelemetry
ENV OTEL_EXPORTER_OTLP_ENDPOINT=localhost:9319
CMD . $HOME/.otel-dotnet-auto/instrument.sh && dotnet Mw-WebApplication.dll
