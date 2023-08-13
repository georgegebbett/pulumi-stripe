import {WebhookEndpoint} from "pulumi-stripe";


const webhook = new WebhookEndpoint("test", {
    connect: false,
    enabledEvents: ["payment_intent.succeeded", "payout.failed"],
    url: "https://test.ghg.pw"
})
