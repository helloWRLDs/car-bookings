const btn = document.getElementById("searchBtn")
const url = 'http://localhost:9090/api/v1/cars?'
let currentUrl = url
const limit = 13
let offset = 0

btn.addEventListener("click", async(e) => {
    e.preventDefault()
    let request = url
    const color = document.getElementById("color").value
    const model = document.getElementById("filter").value.trim()
    request += `color=${color}&model=${model}&limit=${limit}&offset=${offset}`;
    console.log(url)

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
            div.classList.add('grid', 'items', 'bg-green-700', 'px-5', 'py-2', 'relative');
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
        currentUrl = request.slice(0, request.length - 17)
        offset = 0
        console.log("current url: ", currentUrl)
    } catch(error) {
        console.error(error)
    }
})