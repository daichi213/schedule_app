import * as React from 'react';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import ProTip from '../src/ProTip';
import Link from '../src/Link';
import Copyright from '../src/Copyright';
import Header from '../src/blog/Header';
import Checkout from '../src/checkout/Checkout';
import CssBaseline from '@mui/material/CssBaseline';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { blue, green, red, yellow } from '@mui/material/colors';
import { GetServerSideProps } from 'next';
import { Checkbox } from '@mui/material';
import Paper from '@mui/material/Paper';
import TextField from '@mui/material/TextField';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import Grid from '@mui/material/Grid';
import Button from '@mui/material/Button';
import GlobalStyle from "../src/grobalStyles";
import { constants } from 'http2';
import { useState } from 'react';
import axios from 'axios';


export const getServerSideProps : GetServerSideProps = async (context) => {
  const baseUrl : string = process.env.BASE_URL;
  const res = await fetch(baseUrl+ "/");
  // const res = await fetch("http://192.168.210.1:8080" + "/");
  const data = await res.json();
  if (!data) {
    return {
      notFound: true,
    }
  }
  console.log(data);
  return { props: {data}};
}


export default function Index({data}) {

  const baseUrl : string = process.env.BASE_URL;
  const [title, setTitle] = useState("")
  const [content, setContent] = useState("")
  // const handleToAsync : Promise<string> = () => {
  //   return new Promise((resolve) => {
  //     process.env.BASE_URL;
  //   })
  // }

  type todoRes = {
    ID: number
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: string | null
    title: string
    content: string
    status: number
    user_id: number
  }
  
  const headerLinks = [
    {title: 'About', url: "/about"},
    {title: "Create", url: "#"},
    {title: "Create", url: "#"},
    {title: "Create", url: "#"},
    {title: "Create", url: "#"},
    {title: "Create", url: "#"},
  ]
  
  const theme = createTheme({
    typography: {
      fontFamily: "\"RocknRoll One\", \"sans-serif\"",
    },
    palette: {
      primary: {
        main: blue[300],
      },
      secondary: {
        // main: '#bbdefb',
        main: green[200],
      },
    },
  })
  
  const saveTodoTitle = (e) => {
    setTitle(e.target.value);
  }
  
  const saveTodoContent = (e) => {
    setContent(e.target.value);
  }
  
  const handleSubmit = async (event) => {
    // event.preventDefault();
    const resp = await axios({
      method: 'POST',
      url: baseUrl + "/new",
      headers: {
        "Content-Type": "application/json"
      },
      data:{
        "title": title,
        "content": content,
      }
    })
  }
  
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <GlobalStyle />
      <Container maxWidth="lg">
        <Box sx={{ my: 4 }}>
          <Header sections={headerLinks} title={data.page_title}/>
          {/* 子要素のスペースを制御する */}
          <Grid container spacing={3}>
            <Typography variant="h6" component="h1" gutterBottom>
              <List>
                {data.todos.map(todo => 
                  (
                  <Grid item xs={12} sm={12}>
                    <Paper variant="outlined" sx={{ my: { xs: 3, md: 6 }, p: { xs: 2, md: 3 } }}>
                      <ListItem xs={12}>
                        <Link href="#" color="primary">{todo.title}</Link>
                      </ListItem>
                    </Paper>
                  </Grid>
                  )
                )}
              </List>
              <TextField
                required
                id="todoInput"
                name="todoInput"
                label="Please write down ...."
                fullWidth
                autoComplete="given-name"
                variant="standard"
                onChange={saveTodoContent}
              />
              <Button
                color="primary"
                variant="contained"
                onClick={handleSubmit}
              >
                登録
              </Button>
            </Typography>
          </Grid>
          {/* <Checkout /> */}
          {/* <Link href="/about" color="secondary">
            Go to the about page
          </Link> */}
          <ProTip />
          <Copyright />
        </Box>
      </Container>
    </ThemeProvider>
  );
}
