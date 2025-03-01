import htmx from "htmx.org";
import "htmx-ext-head-support";

htmx.logAll();

/** @type {Partial<import("htmx.org").HtmxExtension>} */
const swapAll = {
	onEvent: function (name, event) {
		if (name === "htmx:beforeSwap") {
			/** @type {CustomEvent} */ (event).detail.shouldSwap = true;
		}
		return true;
	},
};

htmx.defineExtension("swap-all", swapAll);
