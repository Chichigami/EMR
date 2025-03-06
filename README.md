# EMR

An electronic medical record website

## Goals

Make a blazingly fast front and backend with the GoTH stack
- Learn front end works
- Learn back end rendering
- Learn Docker + cloud set up
- Deploying the website

# Dependancies

- [Gin](https://gin-gonic.com/)

- [HTMX](https://htmx.org/)

- [TEMPL](https://templ.guide/)

# Current features
- JWT Auth login
- Daily interactive dashboard
- Simple charting (Because charting depends on speciality, I don't have a plan on full support)
- Front end for front desk
- Redis caching for fast dashboards
- PostgreSQL database

# Planned features
- Better css. (It's ugly at the momment 😦)
- Refractoring SQL queries
- Dockerfile

# How to use and deploy EMR yourself
1. Clone the repository.
2. Install [Go](https://go.dev/doc/install).
3. Run `go build`.
4. Run `./EMR`.
