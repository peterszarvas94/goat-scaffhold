// @ts-check

htmx.logAll();

/** @type {import("htmx.org/dist/htmx").HtmxExtension} */
const myExtension = {
  onEvent: function (name, evt) {
    console.log("Fired event: " + name, evt);
  },
};

htmx.defineExtension("my-ext", myExtension);
