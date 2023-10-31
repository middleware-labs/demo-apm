# Real User Monitoring and Browser Traces

To enable Real User Monitoring (RUM) and browser traces, add the following to the `<head>` section of your project

```HTML
<script src="https://cdnjs.middleware.io/browser/libs/0.0.1/middleware-rum.min.js" type="text/javascript"></script>
<script  type="text/javascript">
   window.Middleware &&
    window.Middleware.track({
         projectName:"{APM-PROJECT-NAME}",
         serviceName:"{APM-SERVICE-NAME}",
         accountKey:"{ACCOUNT_KEY}",
         target:"https://{ACCOUNT-UID}.middleware.io"
    });
</script>
```