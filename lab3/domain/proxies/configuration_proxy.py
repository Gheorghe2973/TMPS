from abc import ABC, abstractmethod
from typing import Dict, List
from domain.models.computer import Computer


class ComputerConfigurationInterface(ABC):
    @abstractmethod
    def get_configuration(self, name: str) -> Computer:
        pass
    
    @abstractmethod
    def save_configuration(self, name: str, computer: Computer) -> bool:
        pass
    
    @abstractmethod
    def delete_configuration(self, name: str) -> bool:
        pass


class RealComputerConfiguration(ComputerConfigurationInterface):
    def __init__(self):
        self._configurations: Dict[str, Computer] = {}
    
    def get_configuration(self, name: str) -> Computer:
        if name in self._configurations:
            return self._configurations[name]
        raise ValueError(f"Configuration '{name}' not found")
    
    def save_configuration(self, name: str, computer: Computer) -> bool:
        self._configurations[name] = computer
        return True
    
    def delete_configuration(self, name: str) -> bool:
        if name in self._configurations:
            del self._configurations[name]
            return True
        return False


class ComputerConfigurationProxy(ComputerConfigurationInterface):
    def __init__(self, real_configuration: RealComputerConfiguration):
        self._real_configuration = real_configuration
        self._cache: Dict[str, Computer] = {}
        self._access_log: List[str] = []
        self._authenticated_users: set = set()
        self._current_user: str = None
    
    def authenticate(self, username: str, password: str) -> bool:
        if password == "admin":
            self._authenticated_users.add(username)
            self._current_user = username
            self._log(f"User '{username}' authenticated")
            return True
        return False
    
    def logout(self):
        if self._current_user:
            self._log(f"User '{self._current_user}' logged out")
            self._current_user = None
    
    def get_configuration(self, name: str) -> Computer:
        if not self._check_access():
            raise PermissionError("Authentication required")
        
        if name in self._cache:
            self._log(f"Retrieved '{name}' from cache")
            return self._cache[name]
        
        config = self._real_configuration.get_configuration(name)
        self._cache[name] = config
        self._log(f"Retrieved '{name}' from storage")
        
        return config
    
    def save_configuration(self, name: str, computer: Computer) -> bool:
        if not self._check_access():
            raise PermissionError("Authentication required")
        
        result = self._real_configuration.save_configuration(name, computer)
        
        if result:
            self._cache[name] = computer
            self._log(f"Saved configuration '{name}'")
        
        return result
    
    def delete_configuration(self, name: str) -> bool:
        if not self._check_access():
            raise PermissionError("Authentication required")
        
        result = self._real_configuration.delete_configuration(name)
        
        if result:
            if name in self._cache:
                del self._cache[name]
            self._log(f"Deleted configuration '{name}'")
        
        return result
    
    def get_access_log(self) -> List[str]:
        if not self._check_access():
            raise PermissionError("Authentication required")
        return self._access_log.copy()
    
    def _check_access(self) -> bool:
        return self._current_user in self._authenticated_users
    
    def _log(self, message: str):
        import datetime
        timestamp = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        user = self._current_user or "anonymous"
        log_entry = f"[{timestamp}] {user}: {message}"
        self._access_log.append(log_entry)