const btn = document.getElementById("searchBtn")
const paginationPanel = document.getElementById('paginationPanel')
const url = 'http://localhost:9090/api/v1/cars?'
let currentUrl = url
const limit = 8
let offset = 0
let currentPage = 1
let color = ""
let model = ""
let order = "asc"
let sort = "id"

const loadContent = async() => {
    let request = url
    color = document.getElementById("color").value
    model = document.getElementById("filter").value.trim()
    order = document.getElementById('order').value
    sort = document.getElementById('sort').value
    request += `color=${color}&model=${model}&limit=${limit}&offset=${offset}&sort=${sort}&order=${order}`;

    try {
        let response = await fetch(request, {
            method: "GET",
            mode: "cors"
        })
        if (!response.ok) {
            throw new Error("error")
        }
        const data = await response.json()
        const offers = document.getElementById("offers")
        offers.innerHTML = '';
        data.content.forEach(elem => {
            const div = document.createElement('div');
            div.classList.add('grid', 'items', 'bg-slate-700', 'text-white', 'px-5', 'py-2', 'relative');
            div.innerHTML = `
                <h5 class="text-xl">${elem.vendor} ${elem.model}</h5>
                <p class="text-sm">Year: <span>${elem.year}</span></p>
                <p class="text-sm">Body: <span>${elem.body}</span></p>
                <p class="text-sm">Engine: <span>${elem.engine_capacity}</span>L</p>
                <p class="text-sm">Color: <span>${elem.color}</span></p>
                <p class="text-sm">Mileage: <span>${elem.mileage}</span></p>
                <button class="absolute bottom-0 right-0 text-white bg-blue-500 hover:bg-blue-700 px-7 py-2">Rent</button>
            `;
            offers.appendChild(div);
        });
    } catch(error) {
        console.error(error)
    }
}

btn.addEventListener("click", async(e) => {
    e.preventDefault()
    loadContent()
})

const loadPaginationPanel = async() => {
    paginationPanel.innerHTML = ""
    for (i = 0; i < 5; i++) {
        const btn = document.createElement('button')
        btn.classList.add("px-3", "py-1", "pg-btn")
        if (offset === i * limit) {
            btn.classList.add("bg-blue-500", "text-white")
            btn.disabled = true
        } else {
            btn.classList.add("bg-slate-300", "text-gray-700", "hover:bg-slate-400", "hover:text-white")
        }
        btn.value = i * limit
        btn.innerHTML = i + 1
        paginationPanel.appendChild(btn)
        btn.addEventListener("click", (e) => {
            e.preventDefault()
            btn.disabled = true
            offset = btn.value
            loadContent()
            loadPaginationPanel()
        })
    }
}


window.onload = () => {
    loadPaginationPanel()
    loadContent()
}