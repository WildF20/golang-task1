package category

import "context"

type CategoryChecker interface {
    ExistsByID(ctx context.Context, categoryID string) (bool, error)
}