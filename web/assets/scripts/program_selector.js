/**
 * @type {HTMLSelectElement}
 */
const programInput = document.getElementById("program");

/**
 * @type {HTMLSelectElement}
 */
const majorInput = document.getElementById("major");

/**
 * @typedef {Object} Program
 * @property {string} ID
 * @property {string} Title
 */

/**
 * @typedef {Object} Major
 * @property {string} ID
 * @property {string} Title
 */

/**
 * @typedef {Object} Entry
 * @property {Program} Program
 * @property {Major[]} Majors
 */

/**
 * @type {Entry[]}
 */
const programs = JSON.parse(document.getElementById("programs").textContent);

programInput.addEventListener("change", (e) => {
  /**
   * @type {string}
   */
  const programID = e.target.value;

  const program = programs.filter(
    (program) => program.Program.ID === programID,
  )[0];

  createMajorOptions(program.Majors);
});

/**
 * @param {Major[]} majors
 */
function createMajorOptions(majors) {
  const children = [];

  const option = document.createElement("option");
  option.value = "";
  option.innerText = "Select a major";
  option.selected = true;
  children.push(option);

  for (const major of majors) {
    const option = document.createElement("option");
    option.value = major.ID;
    option.innerText = major.Title;

    children.push(option);
  }

  majorInput.innerHTML = "";
  for (const child of children) {
    majorInput.appendChild(child);
  }
}
