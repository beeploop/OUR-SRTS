let programs = []

const defaultProgram = document.getElementById('defaultProgram')
const defaultMajor = document.getElementById('defaultMajor')
const defaultType = document.getElementById('defaultType')
const defaultCivilStatus = document.getElementById('defaultCivilStatus')

const type = document.getElementById('type')
for (const option of type.options) {
    if (option.value === defaultType.value) {
        option.selected = true;
    }
}

const civilStatus = document.getElementById('civilStatus')
for (const option of civilStatus.options) {
    if (option.value === defaultCivilStatus.value) {
        option.selected = true;
    }
}

document.getElementById('editStudent').addEventListener('click', async () => {
    document.getElementById('modalEditStudent').classList.toggle('hidden')


    if (programs.length > 0) {
        return
    }

    console.log(programs.length)
    console.log("Fetching programs")
    const res = await fetch("/admin/programs")
    if (!res.ok) {
        console.log("Error fetching programs")
        return
    }
    const parsed = await res.json()
    programs = parsed.data

    const selection = document.getElementById('programs')

    programs.forEach(program => {
        const option = document.createElement("option");
        option.value = program.Program;
        option.textContent = program.Program;
        if (program.Program === defaultProgram.value) {
            option.selected = true
        }
        selection.appendChild(option);
    })

    updateMajor()
})
document.getElementById('cancelEditStudent').addEventListener('click', () => {
    document.getElementById('modalEditStudent').classList.toggle('hidden')
})

function updateMajor() {
    const programSelect = document.getElementById("programs");
    const majorSelect = document.getElementById("majors");

    // Get the selected program
    const selectedProgram = programSelect.value;

    // Find the selected program in the array
    const program = programs.find(p => p.Program === selectedProgram);

    // Clear previous options
    majorSelect.innerHTML = '<option value="">Select a Major</option>';

    // Populate majors dropdown if a program is selected
    if (program) {
        program.Majors.forEach(major => {
            const option = document.createElement("option");
            option.value = major;
            option.textContent = major;
            if (major === defaultMajor.value) {
                option.selected = true
            }
            majorSelect.appendChild(option);
        });
    }

    console.log({ program })
}
