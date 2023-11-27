const items = document.getElementById('sidebar-ul')

paths = window.location.href.split('?')[0].split('/')
path = paths.pop()
if (path === '') {
    path = paths.pop()
}

for (const e of items.children) {
    if (e.id === path) {
        e.classList.add('highlight')
    } else {
        e.classList.remove('highlight')
    }
}
