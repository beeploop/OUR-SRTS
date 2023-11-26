const items = document.getElementById('sidebar-ul')

path = window.location.href.split('/').pop().split('?')[0]

for (const e of items.children) {
    if (e.id === path) {
        e.classList.add('highlight')
    } else {
        e.classList.remove('highlight')
    }
}
