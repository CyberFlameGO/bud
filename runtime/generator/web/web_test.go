package web_test

import (
	"context"
	"testing"

	"github.com/livebud/bud/internal/cli"
	"github.com/livebud/bud/internal/cli/testcli"
	"github.com/livebud/bud/internal/is"
	"github.com/livebud/bud/internal/testdir"
)

func TestNoProject(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	dir := t.TempDir()
	td := testdir.New(dir)
	cli := testcli.New(cli.New(dir))
	stdout, stderr, err := cli.Run(ctx)
	is.NoErr(err)
	is.In(stdout.String(), "bud")
	is.Equal(stderr.String(), "")
	is.NoErr(td.NotExists("bud/.app/web/web.go"))
}

func TestEmptyProject(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	dir := t.TempDir()
	td := testdir.New(dir)
	is.NoErr(td.Write(ctx))
	cli := testcli.New(cli.New(dir))
	app, stdout, stderr, err := cli.Start(ctx, "run")
	is.NoErr(err)
	defer app.Close()
	res, err := app.Get("/")
	is.NoErr(err)
	is.Equal(res.Status(), 200)
	is.In(res.Body().String(), "Hey Bud")
	is.NoErr(td.Exists("bud/.app/web/web.go"))
	is.Equal(stdout.String(), "")
	is.Equal(stderr.String(), "")
}
