import React from 'react'
import { Link } from 'react-router-dom'
import { Layout, Menu, Icon } from 'antd'

const { Sider } = Layout

export default class NavSide extends React.Component {
  state = {
    collapsed: false,
    navWidth: 200,
  }

  onCollapse = collapsed => {
    this.setState({ collapsed })
    if (this.state.navWidth === 200) {
      this.setState({ navWidth: 80 })
    } else {
      this.setState({ navWidth: 200 })
    }
  }

  render() {
    return (
      <Layout style={{ maxWidth: this.state.navWidth }}>
        <Sider
          collapsible
          collapsed={this.state.collapsed}
          onCollapse={this.onCollapse}>
          <Menu theme='dark' defaultSelectedKeys={['1']} mode='inline'>
            <Menu.Item className='navlink' key='1'>
              <Icon type='pie-chart' />
              <span>
                <Link className='navlink' to='/'>
                  Home
                </Link>
              </span>
            </Menu.Item>

            <Menu.Item className='navlink' key='2'>
              <span>
                <Icon type='user' />
                <span>
                  <Link className='navlink' to='/investors'>
                    Investors
                  </Link>
                </span>
              </span>
            </Menu.Item>
          </Menu>
        </Sider>
      </Layout>
    )
  }
}
