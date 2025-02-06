// @ts-check
htmx.logAll();

/** @type {import("./htmx/htmx").HtmxExtension} */
const myExtension = {
  onEvent: function (_name, _evt) {
    // ...
  },
};

htmx.defineExtension("my-ext", myExtension);
