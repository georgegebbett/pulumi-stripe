import {WebhookEndpoint} from "@pulumi/stripe";
import {Config} from "@pulumi/pulumi";


const webhook = new WebhookEndpoint("test", {
    connect: false,
    enabledEvents: ["payment_intent.succeeded"],
    url: "https://test.ghg.pw"
})