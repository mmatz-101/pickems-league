import React from "react";
import { useForm } from "@mantine/form";
import {
  Anchor,
  Button,
  Checkbox,
  Container,
  Group,
  Paper,
  PasswordInput,
  Text,
  TextInput,
  Title,
} from "@mantine/core";
import { json } from "react-router-dom";
import { valueGetters } from "@mantine/core/lib/Box/style-system-props/value-getters/value-getters";

function AuthLogin() {
  const form = useForm({
    initialValues: {
      username: "",
      password: "",
      rememberMe: false,
    },
  });

  type FormValues = typeof form.values;

  async function handleLoginForm({ values }) {
    try {
      let resp = await fetch("http://localhost:8000/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(values),
      }).then((r) => r);

      if (resp.status === 200) {
        console.log(resp.json()["data"]);
        // form.reset();
      } else {
        console.log("failure to get request");
      }
    } catch (err) {
      console.log(err);
    }
  }

  return (
    <Container size={420} my={40}>
      <Title
        align="center"
        sx={(theme) => ({
          fontWeight: 900,
        })}
      >
        Welcome Back!
      </Title>
      <Text color="dimmed" size="sm" align="center">
        Do you have an account?{" "}
        <Anchor href="/sign-up" size="sm">
          Create an account.
        </Anchor>
      </Text>

      <Paper withBorder shadow="md" p={30} mt={30} radius="md">
        <form
          onSubmit={form.onSubmit((values: FormValues) =>
            handleLoginForm({ values })
          )}
        >
          <TextInput
            label="Email"
            placeholder="email@email.com"
            required
            {...form.getInputProps("username")}
          />
          <PasswordInput
            label="Password"
            placeholder="your password."
            required
            mt="md"
            {...form.getInputProps("password")}
          />
          <Group position="apart" mt="lg">
            <Checkbox
              label="Remember me"
              sx={{ lineHeight: 1 }}
              {...form.getInputProps("rememberMe", { type: "checkbox" })}
            />
            <Anchor<"a">
              onClick={(event) => event.preventDefault()}
              href="#"
              size="sm"
            >
              Forgot password?
            </Anchor>
          </Group>
          <Button fullWidth mt="xl" type="submit">
            Sign in
          </Button>
        </form>
      </Paper>
    </Container>
  );
}

export default AuthLogin;
