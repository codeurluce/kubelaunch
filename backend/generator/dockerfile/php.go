package dockerfile

import "fmt"

func Laravel(i Input) string {
	return fmt.Sprintf(`
FROM php:8.2-apache

WORKDIR /var/www/html

COPY . .

EXPOSE %d

CMD ["apache2-foreground"]
`, i.Port)
}
