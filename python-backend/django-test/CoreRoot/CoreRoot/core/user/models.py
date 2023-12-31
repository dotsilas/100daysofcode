import uuid

from django.db import models
from django.contrib.auth.models import AbstractBaseUser, PermissionsMixin, BaseUserManager
from django.core.exceptions import ObjectDoesNotExist
from django.http import Http404

# Create your models here.
class UserManager(BaseUserManager):
    def get_object_by_public_id(self, public_id):
        try:
            instance = self.get(public_id=public_id)
            return instance
        except (ObjectDoesNotExist, ValueError, TypeError):
            return Http404
    
    def create_user(self, username, email, password=None, **kwargs):
        """Create and return a 'User' with an email, phone, number, username and password"""
        if username == None:
            raise TypeError("User must have a username")
        if email == None:
            raise TypeError("User must have a email")
        if password is None:
            raise TypeError("user must have a password")

        user = self.model(username=username,
                          email=self.normalize_email(email),
                          **kwargs)

        user.set_password(password)
        user.save(using=self._db)
        return user

    def create_superuser(self, username, email, password, **kwargs):
        """Create and return a 'User' with superuser (admin) permissions"""
        if username == None:
            raise TypeError("Superusers must have a username")
        if email == None:
            raise TypeError("Superusers must have a email")
        if password is None:
            raise TypeError("Superusers must have a password")

        user = self.create_user(username, email, password, **kwargs)
        user.is_superuser = True
        user.is_staff = True
        user.save(using=self._db)

        return user

class User(AbstractBaseUser, PermissionsMixin):
    public_id = models.UUIDField(db_index=True, unique=True, default=uuid.uuid4, editable=False)
    last_name = models.CharField(db_index=True, max_length=255, unique=True)
    first_name = models.CharField(max_length=255)
    username = models.CharField(max_length=255)
    email = models.EmailField(db_index=True, unique=True)
    is_active = models.BooleanField(default=True)
    is_superuser = models.BooleanField(default=False)
    created = models.DateTimeField(auto_now=True)
    updated = models.DateTimeField(auto_now=False)

    USERNAME_FIELD = 'email'
    REQUIRED_FIELDS = ['username']

    objects = UserManager()

    def __str__(self):
        return f"{self.email}"

    @property
    def name(self):
        return f"{self.first_name} {self.last_name}"
