// @ts-check
import { say } from "goaty";

/** @type {string} */
const what = "hello";
console.log(say(what));

htmx.logAll();

/** @type {import("htmx").HtmxExtension} */
const swapAll = {
  onEvent: function (_name, evt) {
    evt.detail.shouldSwap = true;
  },
};

htmx.defineExtension("swap-all", swapAll);
