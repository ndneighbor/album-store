import React, { useState, useEffect } from "react";
import {
  Box,
  Badge,
  Image,
  Input,
  Button,
  FormControl,
  FormLabel,
  FormErrorMessage,
  Flex,
} from "@chakra-ui/react";
import { Formik, Form, Field } from "formik";
import "./App.css";

interface Album {
  id: number;
  title: string;
  price: number;
  artist: string;
  image: string;
}

const AlbumBlock = (album: Album) => {
  return (
    <div key={album.id}>
      <Box
        m="4"
        maxW="sm"
        borderWidth="1px"
        borderRadius="lg"
        overflow="hidden"
      >
        <Box p="6">
          <Box display="flex" alignItems="center">
            <Image
              borderRadius="lg"
              boxSize="100px"
              objectFit="cover"
              src={album.image}
            />
            <Box m={2}>
              <Box display="flex">
                <Badge borderRadius="full" px="2" colorScheme="teal">
                  {album.price}
                </Badge>
                <Box
                  color="gray.500"
                  fontWeight="semibold"
                  letterSpacing="wide"
                  fontSize="xs"
                  textTransform="uppercase"
                  ml="2"
                >
                  {album.artist}
                </Box>
              </Box>

              <Box
                mt="1"
                fontWeight="semibold"
                as="h4"
                lineHeight="tight"
                display="flex"
                textAlign="initial"
              >
                {album.title}
              </Box>
            </Box>
          </Box>
        </Box>
      </Box>
    </div>
  );
};

function App() {
  const [data, setData] = useState([]);
  const [isSubmitting, setSubmitting] = useState(false);

  function fetchData() {
    fetch("http://localhost:4000/albums")
      .then((response) => response.json())
      .then((data) => setData(data));
  }
  useEffect(() => {
    fetchData();
  }, []);

  const albums = data.map((album: Album) => <AlbumBlock {...album} />);

  function validateField(value: string) {
    let error;
    if (!value) {
      error = "Value is required";
    }
    return error;
  }

  return (
    <div className="App">
      <Box bg="teal" w="100%" p="4" color="white">
        Album Store Admin Panel
      </Box>
      <Flex>
        <Box>{albums}</Box>
        <Box w="50%" m="4">
          <Formik
            initialValues={{
              title: "Album Name",
              artist: "Name",
              image: "",
              price: 0,
            }}
            onSubmit={(values, actions) => {
              fetch("http://localhost:4000/albums", {
                method: "POST",
                body: JSON.stringify({ ...values }),
              });
              actions.setSubmitting(false);
              fetchData();
            }}
          >
            {(props) => (
              <Form>
                <Field name="title" validate={validateField}>
                  {({ field, form }) => (
                    <FormControl
                      isInvalid={form.errors.name && form.touched.name}
                    >
                      <FormLabel htmlFor="title">Album Title</FormLabel>
                      <Input {...field} id="title" placeholder="Album Title" />
                      <FormErrorMessage>{form.errors.name}</FormErrorMessage>
                    </FormControl>
                  )}
                </Field>
                <Field name="artist" validate={validateField}>
                  {({ field, form }) => (
                    <FormControl
                      isInvalid={form.errors.name && form.touched.name}
                    >
                      <FormLabel htmlFor="artist">Artist</FormLabel>
                      <Input {...field} id="artist" placeholder="Artist" />
                      <FormErrorMessage>{form.errors.name}</FormErrorMessage>
                    </FormControl>
                  )}
                </Field>
                <Field name="image" validate={validateField}>
                  {({ field, form }) => (
                    <FormControl
                      isInvalid={form.errors.name && form.touched.name}
                    >
                      <FormLabel htmlFor="image">Image</FormLabel>
                      <Input {...field} id="image" placeholder="Album Image" />
                      <FormErrorMessage>{form.errors.name}</FormErrorMessage>
                    </FormControl>
                  )}
                </Field>
                <Field type="number" name="price" validate={validateField}>
                  {({ field, form }) => (
                    <FormControl
                      isInvalid={form.errors.name && form.touched.name}
                    >
                      <FormLabel htmlFor="price">Price</FormLabel>
                      <Input
                        type="number"
                        {...field}
                        id="price"
                        placeholder="Album Price"
                      />
                      <FormErrorMessage>{form.errors.name}</FormErrorMessage>
                    </FormControl>
                  )}
                </Field>

                <Button
                  mt={4}
                  colorScheme="teal"
                  isLoading={props.isSubmitting}
                  disabled={isSubmitting}
                  type="submit"
                >
                  Submit
                </Button>
              </Form>
            )}
          </Formik>
        </Box>
      </Flex>
    </div>
  );
}

export default App;
