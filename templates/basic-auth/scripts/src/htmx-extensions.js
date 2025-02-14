// @ts-check
htmx.logAll();

/** @type {import("htmx").HtmxExtension} */
const swapAll = {
  onEvent: function (_name, evt) {
    evt.detail.shouldSwap = true;
  },
};

htmx.defineExtension("swap-all", swapAll);
