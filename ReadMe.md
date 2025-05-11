# Go Dev

## Useful tools

- [Air](https://github.com/air-verse/air)
    - Air config
        - .air.toml

        ```toml
        root = "."
        testdata_dir = "testdata"
        tmp_dir = "tmp"

        [build]
          args_bin = []
          bin = "./tmp/main"
          cmd = "go build -o ./tmp/main cmd/web/*"
          delay = 1000
          exclude_dir = ["assets", "tmp", "vendor", "testdata", "migrations"]
          exclude_file = []
          exclude_regex = ["_test.go"]
          exclude_unchanged = false
          follow_symlink = false
          full_bin = ""
          include_dir = ["cmd", "internal", "static", "templates"]
          include_ext = ["go", "tpl", "tmpl", "html"]
          include_file = []
          kill_delay = "0s"
          log = "build-errors.log"
          poll = false
          poll_interval = 0
          post_cmd = []
          pre_cmd = []
          rerun = false
          rerun_delay = 500
          send_interrupt = false
          stop_on_error = false

        [color]
          app = ""
          build = "yellow"
          main = "magenta"
          runner = "green"
          watcher = "cyan"

        [log]
          main_only = false
          silent = false
          time = false

        [misc]
          clean_on_exit = false

        [proxy]
          app_port = 0
          enabled = false
          proxy_port = 0

        [screen]
          clear_on_rebuild = false
          keep_scroll = true
        ```

- Soda from the [Buffalo toolset](https://gobuffalo.io)
    - Soda config
        - database.yml

        ```yaml
        development:
          dialect: postgres
          database: db_name
          user: db_user
          password: db_password
          host: ip.add.re.ss
          pool: 5

        test:
          url: {{envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"}}

        production:
          url: {{envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_production"}}
        ```

    - Soda use
        - Examples in migrations folder
        - Create migration files 
            - `soda generate fizz <MigrationName>`
        - Create seed data files
            - `soda generate sql <SeedDataName>`
        - Migrate changes
            - `soda migrate`
        - Revert migration changes
            - `soda migrate down`

## External libraries used

- [Chi v5](https://github.com/go-chi/chi): Application Router
- [NoSurf](https://github.com/justinas/nosurf): CSRF Functionality
- [SCS v2](https://github.com/alexedwards/scs): Session Management
- [goValidator](https://github.com/asaskevich/govalidator): Client side validation
- [godotenv](https://github.com/joho/godotenv): Read .env for application configuration

## External Javascript and CSS

- Bootstrap [v5.3.3](https://getbootstrap.com/docs)
- Popper [v2.11.8](https://popper.js.org/docs)
- DatePicker [v1.3.4](https://mymth.github.io/vanillajs-datepicker)
- Notie [4.3.1](https://github.com/jaredreich/notie)
- SweetAlert2 [v11.15.10](https://sweetalert2.github.io)

## Useful informational videos
- [Udemy GO course](https://www.udemy.com/course/building-modern-web-applications-with-go), coding started at Lecture 21 with keeping of only useful code.  The goal of this project is to provide a web application with persistant data with PostgreSQL.
- [Containerize GO application](https://www.youtube.com/watch?v=1-4hU2e7S4k)
- [Best way to structure GO projects](https://www.youtube.com/watch?v=dxPakeBsgl4)
- [Tool to structure GO projects](https://www.youtube.com/watch?v=1ZbQS6pOlSQ)
    - [Site for video above](https://go-blueprint.dev/)
- LDAP
    - https://dev.to/openlab/ldap-authentication-in-golang-with-bind-and-search-47h5
    - https://cybernetist.com/2020/05/18/getting-started-with-go-ldap/
    - https://usrbinpehli.medium.com/user-authentication-via-ldap-in-go-aff096654db5

## Local .env file

```bash
# is application in production and is template cache used
InProduction: false
# application port number prefixed with ":"
Port_Number: :####
# database name
DB_Name: db_name
# databaase user
DB_User: db_user
# database password
DB_Password: db_password
# database host
DB_Host: ip.add.re.ss
# database port
DB_Port: 5432
# database ssl mode
DB_SSL: disable
# url for ldap server (ldap(s) with 389 or (636))
LDAP_URL: ldaps://ip.add.re.ss:636
# based distringuished name
BASE_DN: dc=home,dc=local
# authenication user (if applicable)
LDAP_ID: ''
# authentication user password (if applicable)
LDAP_PWD: ''
# username identifier
USERNAME: uid
# user organizational unit
USER_OU: ou=people
# group organizational unit
GROUP_OU: ou=groups
# application user group
APP_GROUP_CN: cn=appgroup
```