/* THEME SELECTOR */

const existing = localStorage.getItem("theme");
if (existing) {
	setAttribut(existing);
	setSelection(existing);
}

/** @param {string} value */
function selectTheme(value) {
	localStorage.setItem("theme", value);
	setAttribut(value);
	setSelection(value);
}

/** @param {string} value */
function setAttribut(value) {
	document.body.setAttribute("data-theme", value);
}

/** @param {string} value */
function setSelection(value) {
	const selects = document.querySelectorAll(".goat-theme-selector select");
	selects.forEach((/** @type {HTMLSelectElement} */ element) => {
		element.value = value;
	});
}

/* PASSWORD2 */
const passwordElements = document.querySelectorAll(".goat-password2");
passwordElements.forEach((element) => {
	const input = element.querySelector("input");
	/** @type {HTMLButtonElement} */
	const showButton = element.querySelector(".show-button");
	/** @type {HTMLButtonElement} */
	const hideButton = element.querySelector("button.hide-button");
	if (input && showButton && hideButton) {
		showButton.addEventListener("click", () => {
			input.setAttribute("type", "text");
			showButton.style.display = "none";
			hideButton.style.display = "block";
			hideButton.focus();
		});

		hideButton.addEventListener("click", () => {
			input.setAttribute("type", "password");
			showButton.style.display = "block";
			hideButton.style.display = "none";
			showButton.focus();
		});
	}
});

/**
 * @param {string} query
 * @param {string} item
 * @returns {boolean}
 */
function matches(query, item) {
	let i = 0;
	for (const char of item) {
		if (char.toLowerCase() === query[i]?.toLowerCase()) i++;
		if (i === query.length) return true;
	}
	return false;
}
