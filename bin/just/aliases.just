# Invokes AI to generate a commit.
[group('project')]
save:
  just claude save

# Invokes AI to generate a commit and then pushes commit.
[group('project')]
done:
  just claude save
  git push
