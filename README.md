# API 48 Hours

API 48 Hours is a project designed to challenge me quickly build an API within 48 hours. This project started **Saturday 21/12/2024 11:00 am** and ended **Monday 23/12/2024 10:30 am**. 

## Features

- **Rapid API Development**: Create robust APIs in minimal time.
- **Scalable and Lightweight**: Built with scalability in mind for a variety of use cases.
- **Authentication and Authorization**: Register, Login, Change password and delete account functions and endpoints, using JWT  to generate and validate tokens with a "secret key".
- **World Maps**: The main function of the project is a CRUD that generates world maps (this function comes from a private project I work on). These maps are based on noise, which is placed in a matrix as values representing the height of each coordinate.
The maps can be rendered as images, where each pixel represents a coordinate, and each color corresponds to a height range: blue for sea level, yellow for beaches, green for plains, gray for mountains, white for mountain peaks, etc.
The noise is created using the **Go-Noise library** algorithm, which I then process to create the height matrix. A feature currently in progress in the main project is the generation of biomes, temperatures, climates, etc., all based on height calculations combined with other data matrices also generated with noise. All maps have a water border generated with a mathematical function I designed in GeoGebra.

![WorldMap](/data/example.png)

## Missing Features
The short development time forced me to leave out features such as:
- Parametrization querys
- Unit test
- Documentation through swagger
- Data base optimization

### Technologies Used

- Golang
- Go-Chi framework
- MySql
- Docker & Docker Compose

## Reflections

Despite our best efforts, the project did not meet its original 48-hour completion target. However, this experience provided valuable insights and learning opportunities:

1. **Time Management Challenges**: Planning and execution needed better alignment to meet tight deadlines.
2. **Scope Adjustments**: Certain features required more time than anticipated, highlighting the importance of defining realistic deliverables.
3. **Learning and Growth**: I gained deeper understanding of the technologies used and identified areas for future improvement.

Moving forward, the lessons learned from this project will be instrumental in setting more achievable goals and refining workflows for similar projects.

This project will continue in another repository completing the missing features and adding more.

## Contact

- **Author**: ZeleZ-T (Celeste Caroline López)
- **GitHub**: [github.com/ZeleZ-T](https://github.com/ZeleZ-T)
