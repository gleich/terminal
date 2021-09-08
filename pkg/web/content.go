package web

const HTML = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>ssh</title>
    <link rel="icon" href="https://mattglei.ch/icons/favicon.ico" />
    <link
      rel="stylesheet"
      href="https://cdn.jsdelivr.net/npm/victormono@latest/dist/index.min.css"
    />
  </head>

  <style>
    body {
      background-color: black;
      color: white;
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
        Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
      min-height: 100vh;
      text-align: center;
    }

    h2 {
      font-family: 'Victor Mono';
    }
  </style>
  <body>
    <h1 class="msg">In your terminal emulator of choice:</h1>
    <h2>
      <span style="color: #3ec840">$</span>
      <span style="font-weight: 900; text-decoration: white underline"
        >ssh ssh.mattglei.ch</span
      >
    </h2>
  </body>
</html>`
