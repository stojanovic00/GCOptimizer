// This file can be replaced during build by using the `fileReplacements` array.
// `ng build` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

  const host: string =  "localhost"
  //For this to work you need to start with:  ng serve --host 0.0.0.0 --port 4202
  //This way, it will be accesible from devices in local network
  // const host: string = "192.168.43.111"  

export const environment = {
  production: false,
  applicationPath: "http://" + host + ":8081/application",
  authPath: "http://" + host + ":8081/auth",
  schedulingPath: "http://" + host + ":8082/scheduling",
  scoringPath: "http://" + host + ":8082/scoring",
  webSocketPath: "ws://" + host + ":8082/scoring/web-socket",
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/plugins/zone-error';  // Included with Angular CLI.
