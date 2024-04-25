# The Daily Brief

The Daily Breif, provides a way of glimpsing through, the important healiness, and only watch the important stuff.

## Architecture

The backend server is made with Go and the frontend is made with Sveltekit (JsDocs)

## Contributing

First fork this project and then clone the repo

```bash
git clone <this-repo>
```

Then you can run the server by opening a terminal, and then using the using the following command

```bash
cd server
go run .
```

Your server on port `8080`

You can run the client side, by opening a seperate terminal and then using the following command

```bash
cd client
npm run dev
```

Your client will serve on port `5173`

## Testing

You can test the server coding using the following command

```bash
go run test
```

You can run `unit tests` client side using the following command

```bash
npm run test
```

Testing is done with `vitest`

