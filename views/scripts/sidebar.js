const items = document.getElementById('sidebar-ul')

path = window.location.href.split('/').pop()

for (const e of items.children) {
    if (e.id === path) {
        e.classList.add('highlight')
    } else {
        e.classList.remove('highlight')
    }
}
