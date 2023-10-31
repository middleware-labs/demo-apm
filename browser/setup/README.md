# Real User Monitoring and Browser Traces

To enable Real User Monitoring (RUM) and browser traces, add the following to the `<head>` section of your project.

```HTML
<script src="https://cdnjs.middleware.io/browser/libs/0.0.1/middleware-rum.min.js" type="text/javascript"></script>
<script>
    if (window.Middleware){
        Middleware.track({
            serviceName: "service_name",
            projectName: "project_name",
            accountKey: "account_key",
            target:"https://<UID>.middleware.io",
        })
    }
</script>
```