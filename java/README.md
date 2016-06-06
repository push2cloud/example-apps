# Java example application

This brief example demonstrastes how to deploy a java application with push2cloud.

Using the `path` variable, you can tell push2cloud where your built jar artifact resides. This mirrors the function of the [`path` attribute](http://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html#path) in traditional CloudFoundry manifests.

Using the `scripts`/`package` declaration you can instruct push2cloud to build the artifacts on demand during deployment. Note that the scripts get executed in the same directory as your application manifest resides.

## Example application manifest
Full example can be found [apps/java/application.json](here).
```
{
  "name": "cf-java-example-app",
  "version": "1.0.0",
  "deployment": {
    "path": "build/libs/java-1.0.jar",
    "scripts": {
      "package": [
        "gradle jar"
      ]
    }
  }
}
```
