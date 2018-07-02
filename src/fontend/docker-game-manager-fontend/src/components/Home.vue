<template>
  <div class="hero">
    <b-jumbotron fluid>
      <template slot="header">
        RatNest Server Manager
      </template>
      <template slot="lead">
        Spin up game servers with ease
      </template>
      <hr class="my-4">
      <b-container fluid>
        <b-row>
          <b-col>
            <b-btn :size="'lg'" variant="primary" href="https://discord.gg/xthJVWR">Click here to join the Discord for server information</b-btn>
          </b-col>
        </b-row>
      </b-container>
    </b-jumbotron>
  </div>
</template>

<style>
  .hero {
    height: 90vh;
    display: flex;
    align-items: center;
    justify-content: center;
    text-align: center;
  }
  .hero .lead {
    font-weight: 200;
    font-size: 1.5rem;
  }
</style>