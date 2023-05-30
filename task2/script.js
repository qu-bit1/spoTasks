function fetchAndDisplayData() {

  const url = 'https://swapi.dev/api/planets/1/';

  const cachedData = localStorage.getItem('swapiPlanets');

  if (cachedData) {
    // If data is already cached, parse and display it
    const planetsData = JSON.parse(cachedData);
    displayPlanetData(planetsData);
  } else {
    // If data is not cached, fetch it from the API
  fetch(url)
    .then(response => response.json())
    .then(data => {
      // Cache the fetched data
      localStorage.setItem('swapiPlanets', JSON.stringify(data));
      displayPlanetData(data);
    })
      
      
    .catch(error => {
      console.log('Error:', error);
    });

  }
}

function displayPlanetData(data){
  const planetDataElement = document.getElementById('planet-data');
          planetDataElement.innerHTML = `
            <p>Name: ${data.name}</p>
            <p>Rotation Period: ${data.rotation_period}</p>
            <p>Orbital Period: ${data.orbital_period}</p>
            <p>Diameter: ${data.diameter}</p>
            <p>Climate: ${data.climate}</p>
            <p>Gravity: ${data.gravity}</p>
            <p>Terrain: ${data.terrain}</p>
            <p>Surface Water: ${data.surface_water}</p>
            <p>Population: ${data.population}</p>
          `;

          // Fetching residents' names
          const residentsListElement = document.getElementById('residents-list');
          data.residents.forEach(residentUrl => {
            fetch(residentUrl)
              .then(response => response.json())
              .then(residentData => {
                const residentName = residentData.name;
                const residentListItem = document.createElement('li');
                residentListItem.textContent = residentName;
                residentsListElement.appendChild(residentListItem);
              })
              .catch(error => {
                console.log('Error:', error);
              });
          });

          // Fetching film titles
          const filmElement = document.getElementById('films-list');
          data.films.forEach(filmUrl => {
            fetch(filmUrl)
              .then(response => response.json())
              .then(filmData => {
                const filmTitle = filmData.title;
                const filmsListItem = document.createElement('li');
                filmsListItem.textContent = filmTitle;
                filmElement.appendChild(filmsListItem);
              })
              .catch(error => {
                console.log('Error:', error);
              })
          });
      }

      

fetchAndDisplayData();