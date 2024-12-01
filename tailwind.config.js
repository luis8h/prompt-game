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
                primary: "#001c2a",
                secondary: "#2f4757",
                highlight: "#cdf6ff",
            }
        },
    },
    plugins: [
    ]
}

