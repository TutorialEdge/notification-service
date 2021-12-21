import * as sst from "@serverless-stack/resources";

export default class MyStack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);

    // Create a HTTP API
    const api = new sst.Api(this, "Api", {
      defaultFunctionProps: {
        timeout: 30,
      },
      environment: {
        DB_HOST: process.env.DB_HOST,
        DB_NAME: process.env.DB_NAME,
        DB_USERNAME: process.env.DB_USERNAME,
        DB_PASSWORD: process.env.DB_PASSWORD,
        DB_PORT: process.env.DB_PORT,
        DB_SSL: process.env.DB_SSL,
        MAILERLITE_API_KEY: process.env.MAILERLITE_API_KEY,
      },
      routes: {
        "GET /api/v1/notifications/{id}": "cmd/endpoints/notifications/read",
        "POST /api/v1/notifications": "cmd/endpoints/notifications/create",
        "PUT /api/v1/notifications/{id}": "cmd/endpoints/notifications/update",
        "DELETE /api/v1/notifications/{id}": "cmd/endpoints/notifications/delete",
      }
    });

    const queue = new sst.Queue(this, "notificationsQueue", {
      consumer: "cmd/engine/process",
    })

    const cron = new sst.Cron(this, "Processor", {
      environment: {
        queueUrl: queue.sqsQueue.queueUrl,
        DB_HOST: process.env.DB_HOST,
        DB_NAME: process.env.DB_NAME,
        DB_USERNAME: process.env.DB_USERNAME,
        DB_PASSWORD: process.env.DB_PASSWORD,
        DB_PORT: process.env.DB_PORT,
        DB_SSL: process.env.DB_SSL,
      },
      schedule: "rate(1 minute)",
      job: "cmd/engine/schedule"
    })

    cron.attachPermissions([queue]);

    // Show the endpoint in the output
    this.addOutputs({
      "ApiEndpoint": api.url,
      "QueueURL": queue.sqsQueue.queueUrl,
    });
  }
}
