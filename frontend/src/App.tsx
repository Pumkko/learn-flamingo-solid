import { Match, Switch, createSignal } from 'solid-js'
import solidLogo from './assets/solid.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { createQuery } from '@tanstack/solid-query'

function fetchGreeting() {
  return fetch("/api/greet/pikachu").then(f => f.json())
}

function App() {
  const [count, setCount] = createSignal(0)


  const query = createQuery(() => ({
    queryKey: ['greeting'],
    queryFn: fetchGreeting,
  }))

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} class="logo" alt="Vite logo" />
        </a>
        <a href="https://solidjs.com" target="_blank">
          <img src={solidLogo} class="logo solid" alt="Solid logo" />
        </a>
      </div>
      <h1>Vite + Solid</h1>
      <div>
        <Switch>
          <Match when={query.isPending}>
            <p>Loading...</p>
          </Match>
          <Match when={query.isError}>
            <p>Error: {query.error?.message}</p>
          </Match>
          <Match when={query.isSuccess}>
            <p>Hello {query.data.Nickname}</p>
          </Match>
        </Switch>
      </div>
      <div class="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count()}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p class="read-the-docs">
        Click on the Vite and Solid logos to learn more
      </p>
    </>
  )
}

export default App
