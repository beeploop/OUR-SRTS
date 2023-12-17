/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./views/templates/**/*.tmpl"],
    theme: {
        extend: {
            fontFamily: {
                quicksand: ["Quicksand", "sans-serif"],
            },
        },
    },
    plugins: [],
}

