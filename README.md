![](./Six910Cart-trans-blue.png)

# Pronounced: Six 9 10

## [Swagger API Docs](http://api-swagger-docs.six910.com/swagger/index.html)

[![Build Status](https://travis-ci.org/Ulbora/Six910.svg?branch=master)](https://travis-ci.org/Ulbora/Six910)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Six910&metric=alert_status)](https://sonarcloud.io/dashboard?id=Six910)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/Six910)](https://goreportcard.com/report/github.com/Ulbora/Six910)


## Bitcoin payments
Six910 now has a bitcoin payment module built on open source BTCPay Server. BTCPay Server is a self-hosted, open-source cryptocurrency payment processor. It's secure, private, censorship-resistant and free.


## Docker
https://hub.docker.com/r/ulboralabs/six910

### Project Status
Six910 has released and currently running on multiple sites.


### The Speeding Fast Shopping Cart ECommerce System with Dependency Injection
An open source Ecommerce platform server for building online shopping cart systems and now has a cryptocurrency payment module built on BTCPay Server among others; it is written in golang. This project is the REST service backend. The UI project will including an admin panel and switchable templates.

All components of Six910 are developed as Go modules and injected using dependency injection.



This project is the REST implementation of the Six910 Ecommerce solution. 

1. Users of the Six910, together with Six910 UI are able to customize templates just like other hosted solutions.
2. Templates can either be written in Golang templating or use a JavaScript framework like Angular or React.
3. REST services expose all cart functionality.

## Template Designer
There will also be a template designer to make desiging templates much easier than it currently is with most hosted shopping cart solutions.

1. Users can download there current template
2. Modify the template using the user's store URL
3. Upload the template back to the hosted site

## Database
### The database module in injected using dependency injection in the main func.
The default database for Six910 is MySQL. The database module can be found [here](https://github.com/Ulbora/six910-mysql). The database interface, however is modular and can easily be switched out for any other database.

## Addon Applications
Six910 allows third party developers to build addon applications that integrate into templates.



