const url = 'http://localhost:9090'

indexRender = async() => {
    try {
        let response = await fetch(url + '/api/v1/cars?limit=8&offset=0', {
            method: "GET",
            mode: "cors"
        })
        if (!response.ok) {
            throw new Error("error")
        }
        const data = await response.json()
        const offers = document.getElementById("offers")
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
    } catch(error) {
        console.error(error)
    }
}

window.onload = () => {
    indexRender()
}