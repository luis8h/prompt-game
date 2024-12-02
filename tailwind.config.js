const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        'views/**/*.templ',
    ],
    theme: {
        container: {
            center: true,
            padding: {
                DEFAULT: "1rem",
                mobile: "2rem",
                tablet: "4rem",
                desktop: "5rem",
            },
        },
        extend: {
            colors: {
                primary: "#00111a",
                secondary: "#002233",
                highlight: "#75b8c7",
                textcolor: "#f2ddcc",
            }
        },
    },
    plugins: [
    ]
}

