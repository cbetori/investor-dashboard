import React from 'react'

import AppContainer from './AppContainer'
import middleware from './middleware'
import { Provider } from 'react-redux'
import { BrowserRouter } from 'react-router-dom'

const Store = () => {
  return (
    <div id='store'>
      <Provider store={middleware}>
        <BrowserRouter>
          <AppContainer />
        </BrowserRouter>
      </Provider>
    </div>
  )
}

export default Store
