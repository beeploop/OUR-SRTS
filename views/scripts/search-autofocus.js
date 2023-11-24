const searchbar = document.getElementById('search-bar')
const urlParams = new URLSearchParams(window.location.search);
const term = urlParams.get('term');

console.log(term)
if (term) {
    searchbar.value = term
}

